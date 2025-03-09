package facade

import "fmt"

type MediaFacade struct {
	Audio        *AudioSystem
	Video        *VideoSystem
	Notification *NotificationSystem
}

func (m *MediaFacade) Play() {
	fmt.Println("[Facade] >>> Play video")
	m.Audio.setSpeed(1)
	m.Audio.setVolume(60)
	m.Video.setSpeed(1)
	m.Video.setResolution(1440, 1440)
}

func NewMediaFacade() *MediaFacade {
	return &MediaFacade{
		Audio:        newAudioSystem(),
		Video:        newVideoSystem(),
		Notification: newNotificationSystem(),
	}
}

type AudioSystem struct {
	volume, speed int
}

func (a *AudioSystem) setSpeed(speed int) {
	fmt.Printf("Audio Speed set to %d\n", speed)
	a.speed = speed
}

func (a *AudioSystem) setVolume(volume int) {
	fmt.Printf("Audio Volume set to %d\n", volume)
	a.volume = volume
}

func newAudioSystem() *AudioSystem {
	return &AudioSystem{
		volume: 50,
		speed:  1,
	}
}

type VideoSystem struct {
	speed, width, height int
}

func (v *VideoSystem) setSpeed(speed int) {
	fmt.Printf("Video Speed set to %d\n", speed)
	v.speed = speed
}

func (v *VideoSystem) setResolution(width int, height int) {
	fmt.Printf("Video Resolution set to %d, %d\n", width, height)
	v.width = width
	v.height = height
}

func newVideoSystem() *VideoSystem {
	return &VideoSystem{
		speed:  1,
		width:  1024,
		height: 1024,
	}
}

type NotificationSystem struct {
	receivedNotification bool
}

func (n *NotificationSystem) setNotification(notification bool) {
	fmt.Printf("Notification set to %d\n", notification)
	n.receivedNotification = notification
}

func newNotificationSystem() *NotificationSystem {
	return &NotificationSystem{
		receivedNotification: true,
	}
}
