// abstract class APIServer {
//     abstract request(endpoint: string): void
// }

export interface APIService {
    request(endpoint: string): void
}

class RealAPI implements APIService {
    request(endpoint: string): void {
        console.log(`Fetching data from ${endpoint}`)
    }
}

export class APIProxy implements APIService {
    private realAPI: RealAPI
    private requestTimes: Map<string, number>

    constructor() {
        this.realAPI = new RealAPI
        this.requestTimes = new Map()
    }

    request(endpoint: string): void {
        const now = Date.now()
        const lastRequestTime = this.requestTimes.get(endpoint) || 0

        if (now - lastRequestTime < 3000) {
            console.log(`Request to ${endpoint} is blocked（Too many requests）`)
            return
        }

        this.realAPI.request(endpoint)
        this.requestTimes.set(endpoint, now)
    }
}