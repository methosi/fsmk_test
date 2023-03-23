package playlist

import "testing"

func TestPlayList(t *testing.T) {
	// Create new playlist
	p := NewPlaylist(10)

	// Add song to the playlist
	p.AddSong("song1", 1)
	p.AddSong("song2", 1)
	p.AddSong("song3", 1)
	p.AddSong("song4", 1)

	// Test GetSongList
	expectedList := []string{"song1", "song2", "song3", "song4"}
	songList := p.GetSongList()
	if len(songList) != len(expectedList) {
		t.Errorf("Expected %v, got %v", expectedList, songList)
	} else {
		for i, v := range expectedList {
			if songList[i] != v {
				t.Errorf("Expected %v, got %v", expectedList, songList)
			}
		}
	}

	// Test PlaySong with a new song
	p.PlaySong("song5")
	expectedList = []string{"song5", "song1", "song2", "song3", "song4"}
	songList = p.GetSongList()
	if len(songList) != len(expectedList) {
		t.Errorf("Expected %v, got %v", expectedList, songList)
	} else {
		for i, v := range expectedList {
			if songList[i] != v {
				t.Errorf("Expected %v, got %v", expectedList, songList)
			}
		}
	}

	// Test PlaySong with an existing song
	p.PlaySong("song3")
	expectedList = []string{"song3", "song5", "song1", "song2", "song4"}
	songList = p.GetSongList()
	if len(songList) != len(expectedList) {
		t.Errorf("Expected %v, got %v", expectedList, songList)
	} else {
		for i, v := range expectedList {
			if songList[i] != v {
				t.Errorf("Expected %v, got %v", expectedList, songList)
			}
		}
	}
}
