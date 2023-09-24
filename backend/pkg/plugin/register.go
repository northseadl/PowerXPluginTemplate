package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const pluginUri = "plugin/v1/plugins"

func RegisterPluginToHost(host string, plugin *Plugin) error {
	pluginData, _ := json.Marshal(plugin.Data())
	body := bytes.NewBuffer(pluginData)
	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s/%s", host, pluginUri), body)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("register plugin failed, status code %v, body %s", response.StatusCode, string(body))
	}

	return nil
}
