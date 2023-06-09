package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"telegrambot/clients/telegram/lib/e"

)



type Client struct {
	host 		string
	basePath 	string
	clent 		http.Client
}

const (
	getUpdatesMethod = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host string, token string) Client {
	return Client{
		host: host,
		basePath: newBasePath(token),
		clent: http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c  *Client) Update(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("linit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	} 
	
	return res.Result, nil
}

func (c *Client) sendMessage(chatID int, text string) error {
	q:= url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)

	if err != nil {
		return e.Wrap("cannot send message", err)
	}
	return nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	
	const errMsg = "can't do request"
	defer func () {
		err = e.WrapIFErr(errMsg, err)
	}()

	u := url.URL {
		Scheme: "http",
		Host: c.host,
		Path: path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.clent.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() {_ = resp.Body.Close()}()
	
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err 
	}

	return body, nil
}


func (c *Client) SendMessage() {

}