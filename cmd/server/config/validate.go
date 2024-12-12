package config

import (
	"errors"
	"net/url"
	"slices"
)

func validate(conf Config) error {
	if conf.MQTT.Listen == nil {
		return errors.New("mqtt.listen is required")
	}

	for _, row := range conf.MQTT.Listen {
		rowUrl, err := url.Parse(row)
		if err != nil {
			return err
		}

		if !slices.Contains([]string{"mqtt", "ws"}, rowUrl.Scheme) {
			return errors.New("mqtt.listen: invalid scheme")
		}

		if rowUrl.Port() == "" {
			return errors.New("mqtt.listen: port is required")
		}
	}

	return nil
}
