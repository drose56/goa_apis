// Code generated by goa v3.7.6, DO NOT EDIT.
//
// poke HTTP client encoders and decoders
//
// Command:
// $ goa gen poke/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	poke "poke/gen/poke"

	goahttp "goa.design/goa/v3/http"
)

// BuildPokemonRequest instantiates a HTTP request object with method and path
// set to call the "poke" service "pokemon" endpoint
func (c *Client) BuildPokemonRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		name string
	)
	{
		p, ok := v.(*poke.PokemonPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("poke", "pokemon", "*poke.PokemonPayload", v)
		}
		name = p.Name
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PokemonPokePath(name)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("poke", "pokemon", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodePokemonResponse returns a decoder for responses returned by the poke
// pokemon endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodePokemonResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("poke", "pokemon", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("poke", "pokemon", resp.StatusCode, string(body))
		}
	}
}
