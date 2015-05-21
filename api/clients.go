package main

type Client struct {
	Host string
}

func (c *Client) CreateReport() (Report, error) {
	var respTodo Report
	report := Report{}

	r, err := makeRequest(c, "POST", "/reports/", report)
	if err != nil {
		return respTodo, err
	}
	err = processResponseEntity(r, &respTodo, 201)
	return respTodo, err
}
