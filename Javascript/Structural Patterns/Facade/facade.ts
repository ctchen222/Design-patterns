class AudioSystem {
    volume: number
    speed: number
    
    constructor() {
        this.volume = 50
        this.speed = 1
    }

    setSpeed(level: number) {
        this.speed = level
        console.log(`[Audio] >>> Speed set to ${level}`)
    }

    setVolume(level: number) {
        this.volume = level
        console.log(`[Audio] >>> Volume set to ${level}`)
    }
}

class VideoSystem {
    speed: number
    width: number
    height: number

    constructor() {
        this.speed = 1
        this.width = 1024
        this.height = 1024
    }

    setSpeed(level: number) {
        this.speed = level
        console.log(`[Video] >>> Speed set to ${level}`)
    }

    setResolution(width: number, height: number) {
        this.width = width
        this.height = height
        console.log(`[Video] >>> Resolution set to ${width}x${height}`)
    }
}

class NotificationSystem {
    receivedNotification: boolean

    constructor() {
        this.receivedNotification = true
    }

    setNotification(state: boolean) {
        this.receivedNotification = state
        console.log(`[Notification] >>> Notification state set to ${state}`)
    }
}

export class MediaFacade {
    private audio: AudioSystem
    private video: VideoSystem
    private notification: NotificationSystem

    constructor() {
        this.audio = new AudioSystem()
        this.video = new VideoSystem()
        this.notification = new NotificationSystem()
    }

    play() {
        console.log("[Facade] >>> Play video")
        this.audio.setSpeed(1)
        this.audio.setVolume(60)
        this.video.setSpeed(1)
        this.video.setResolution(1440, 1440)
    }

    turnOffNotification() {
        console.log("[Facade] >>> Turn off notification")
        this.notification.setNotification(false)
    }
}