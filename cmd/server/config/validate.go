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
		rowURL, err := url.Parse(row)
		if err != nil {
			return err
		}

		if !slices.Contains([]string{"mqtt", "ws"}, rowURL.Scheme) {
			return errors.New("mqtt.listen: invalid scheme")
		}

		if rowURL.Port() == "" {
			return errors.New("mqtt.listen: port is required")
		}
	}

	return nil
}
