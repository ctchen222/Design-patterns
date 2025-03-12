export interface Service {
    request(user: string): void
}

class RealService implements Service {
    request(user: string): void {
        console.log(`RealService: Handle ${user} request`)
    }
}

export class ProxyService implements Service {
    private realService: RealService
    private allowUsers: Set<string>

    constructor() {
        this.realService = new RealService()
        this.allowUsers = new Set(["admin", "user1"])
    }

    request(user: string): void {
       if (!this.allowUsers.has(user)) {
            console.log(`ProxyService: Reject ${user} request`)
        }
       this.realService.request(user)
    }
}
