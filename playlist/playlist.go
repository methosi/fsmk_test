package playlist

type PlayListInterface interface {
	PlaySong(songTitle string)
	AddSong(songTitle string, value byte)
}

type Playlist struct {
	Song map[string]byte
	Key  []string
}

func NewPlaylist(noOfSong int) *Playlist {
	return &Playlist{
		Song: make(map[string]byte),
		Key:  make([]string, 0, noOfSong),
	}
}

func (p *Playlist) PlaySong(songTitle string) {
	// Check if song is in the playlist
	_, ok := p.Song[songTitle]
	if !ok {
		// If song is not found, add song to the playlist
		p.AddSong(songTitle, 1)
		p.MoveToFirst(songTitle)
	}

	// If song is found, move the song to the first
	p.MoveToFirst(songTitle)
}

// AddSong will add song to the playlist
func (p *Playlist) AddSong(songTitle string, value byte) {
	p.Song[songTitle] = value
	p.Key = append(p.Key, songTitle)
}

func (p *Playlist) GetSongList() []string {
	return p.Key
}

func (p *Playlist) MoveToFirst(songTitle string) {
	idx := -1
	for i, k := range p.Key {
		if k == songTitle {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}

	// Remove the key from the slice
	p.Key = append(p.Key[:idx], p.Key[idx+1:]...)

	// Add the key to the front of the slice
	p.Key = append([]string{songTitle}, p.Key...)
}
