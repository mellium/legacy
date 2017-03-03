// Copyright 2016 Sam Whited.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

// Package compress implements XEP-0138: Stream Compression and XEP-0229: Stream
// Compression with LZW.
package compress

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"

	"mellium.im/xmpp"
	"mellium.im/xmpp/streamerror"
)

// Namespaces used by stream compression.
const (
	NSFeatures = "http://jabber.org/features/compress"
	NSProtocol = "http://jabber.org/protocol/compress"
)

var (
	errNoMethods = errors.New("No supported compression methods were found")
)

// New returns a new xmpp.StreamFeature that can be used to negotiate stream
// compression.
// The returned stream feature always supports ZLIB compression; other
// compression methods are optional.
func New(methods ...Method) xmpp.StreamFeature {
	// TODO: Throw them into a map to dedup and then iterate over that?
	methods = append(methods, zlibMethod)
	return xmpp.StreamFeature{
		Name:      xml.Name{Local: "compression", Space: NSFeatures},
		Necessary: xmpp.Secure | xmpp.Authn,
		List: func(ctx context.Context, e *xml.Encoder, start xml.StartElement) (req bool, err error) {
			if err = e.EncodeToken(start); err != nil {
				return
			}

			methodStart := xml.StartElement{Name: xml.Name{Local: "method"}}

			for _, m := range methods {
				select {
				case <-ctx.Done():
					return req, ctx.Err()
				default:
				}

				if err = e.EncodeToken(methodStart); err != nil {
					return
				}
				if err = e.EncodeToken(xml.CharData(m.Name)); err != nil {
					return
				}
				if err = e.EncodeToken(methodStart.End()); err != nil {
					return
				}
			}

			if err = e.EncodeToken(start.End()); err != nil {
				return
			}
			return false, e.Flush()
		},
		Parse: func(ctx context.Context, d *xml.Decoder, start *xml.StartElement) (bool, interface{}, error) {
			listed := struct {
				XMLName xml.Name `xml:"http://jabber.org/features/compress compression"`
				Methods []string `xml:"http://jabber.org/features/compress method"`
			}{}
			if err := d.DecodeElement(&listed, start); err != nil {
				return false, nil, err
			}

			// TODO: Is this actually an error, or is this a bad implementation that
			//       we can ignore? If someone strips all security mechanisms do we
			//       care that we'll effectively negotiate "none" compression?
			if len(listed.Methods) == 0 {
				return false, nil, errNoMethods
			}

			return true, listed.Methods, nil
		},
		Negotiate: func(ctx context.Context, session *xmpp.Session, data interface{}) (mask xmpp.SessionState, rw io.ReadWriter, err error) {
			conn := session.Conn()

			// If we're a server.
			if (session.State() & xmpp.Received) == xmpp.Received {
				clientSelected := struct {
					XMLName xml.Name `xml:"http://jabber.org/protocol/compress compress"`
					Method  string   `xml:"method"`
				}{}
				if err = session.Decoder().Decode(&clientSelected); err != nil {
					return
				}

				// If no method was selected, or something weird happened with decoding…
				if clientSelected.Method == "" {
					_, err = fmt.Fprint(conn, `<failure xmlns='`+NSProtocol+`'><setup-failed/></failure>`)
					return
				}

				var selected Method
				for _, method := range methods {
					if method.Name == clientSelected.Method {
						selected = method
						break
					}
				}

				// The client requested a method that we did not send…
				if selected.Name == "" {
					_, err = fmt.Fprint(conn, `<failure xmlns='`+NSProtocol+`'><unsupported-method/></failure>`)
					return
				}

				if _, err = fmt.Fprint(conn, `<compressed xmlns='`+NSProtocol+`'/>`); err != nil {
					return
				}

				rw, err = selected.Wrapper(conn)
				return mask, rw, err
			}

			var selected Method
		selectmethod:
			for _, m := range methods {
				for _, name := range data.([]string) {
					if name == m.Name {
						selected = m
						break selectmethod
					}
				}
			}

			if selected.Name == "" {
				return mask, nil, errors.New(`No matching compression method found`)
			}

			_, err = fmt.Fprintf(conn, `<compress xmlns='`+NSProtocol+`'><method>%s</method></compress>`, selected.Name)
			if err != nil {
				return
			}

			d := session.Decoder()
			tok, err := d.Token()
			if err != nil {
				return mask, nil, err
			}

			if t, ok := tok.(xml.StartElement); ok && t.Name.Local == "compressed" && t.Name.Space == NSProtocol {
				if err = d.Skip(); err != nil {
					return mask, nil, err
				}
				rw, err = selected.Wrapper(conn)
				return mask, rw, err
			}

			// TODO: Use appropriate errors.
			return mask, nil, streamerror.BadFormat
		},
	}
}
