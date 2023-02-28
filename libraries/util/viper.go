package util

import (
	"fmt"

	"github.com/spf13/viper"
)

func Viper() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	viperSetting := viper.AllSettings()
	fmt.Println("viperSetting", viperSetting)

}
