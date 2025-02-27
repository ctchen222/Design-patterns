interface Observer {
    update(news: string): void
}

interface Subject {
    attach(observer: Observer): void
    detach(observer: Observer): void
    notify(): void
}

export class NewsAgency implements Subject {
    private observers: Observer[] = []
    private latesNews: string = ""
    
    attach(observer: Observer): void {
        this.observers.push(observer)
    }

    detach(observer: Observer): void {
        this.observers = this.observers.filter(obs => obs !== observer)
    }

    setNews(news: string): void {
        this.latesNews = news
        this.notify()
    }

    notify(): void {
        this.observers.forEach(observer => observer.update(this.latesNews))
    }
}

export class Subscriber implements Observer {
    private name: string

    constructor(name: string) {
        this.name = name
    }

    update(news: string): void {
       console.log(`${this.name} received latest new: ${news}`) 
    }
}