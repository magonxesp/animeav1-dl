// AnimeAV1-DL - Un programa para extraer enlaces de descarga de animeav1.com
// Copyright (C) 2025  MagonxESP
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package common

import (
	"reflect"
	"strings"
	"testing"
)

func TestExtractEpisodesLinks(t *testing.T) {
	expectedLinks := []string{
		"https://animeav1.com/media/yofukashi-no-uta/1",
		"https://animeav1.com/media/yofukashi-no-uta/2",
		"https://animeav1.com/media/yofukashi-no-uta/3",
		"https://animeav1.com/media/yofukashi-no-uta/4",
		"https://animeav1.com/media/yofukashi-no-uta/5",
		"https://animeav1.com/media/yofukashi-no-uta/6",
		"https://animeav1.com/media/yofukashi-no-uta/7",
		"https://animeav1.com/media/yofukashi-no-uta/8",
		"https://animeav1.com/media/yofukashi-no-uta/9",
		"https://animeav1.com/media/yofukashi-no-uta/10",
		"https://animeav1.com/media/yofukashi-no-uta/11",
		"https://animeav1.com/media/yofukashi-no-uta/12",
		"https://animeav1.com/media/yofukashi-no-uta/13",
	}

	url := "https://animeav1.com/media/yofukashi-no-uta"
	links, err := ExtractEpisodesLinks(url)

	if err != nil {
		t.Errorf("ExtractEpisodesLinks returned an error: %v", err)
	}

	if !reflect.DeepEqual(links, expectedLinks) {
		t.Errorf("ExtractEpisodesLinks returned incorrect links.\nExpected: %v\nGot: %v", expectedLinks, links)
	}
}

func TestExtractEpisodeDownloadLink(t *testing.T) {
	url := "https://animeav1.com/media/yofukashi-no-uta/2"
	downloadLink, err := ExtractEpisodeDownloadLink(url)

	if err != nil {
		t.Errorf("ExtractEpisodeDownloadLink returned an error: %v", err)
	}

	if !strings.HasPrefix(downloadLink, "https://mega.nz/") {
		t.Errorf("ExtractEpisodeDownloadLink returned incorrect link. Expected Mega link, got: %v", downloadLink)
	}
}
