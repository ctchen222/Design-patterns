// abstract class HttpClient {
//     abstract get(url: string): Promise<string>
// }

interface HttpClient {
    get(url: string): Promise<string>
}

export class OldHttpClient {
    request(config: { url: string, method: string }): Promise<{ data: string }> {
        return new Promise((resolve) => {
            console.log(`Send request via old http client to ${config.url}`)
            resolve({data: "data from old http request"})
        })
    }
}

export class HttpClientAdaptor implements HttpClient {
    oldHttpClient: OldHttpClient

    constructor(oldHttpClient: OldHttpClient) {
        this.oldHttpClient = oldHttpClient
    }

    async get(url: string): Promise<string> {
        console.log("[Adaptor] Transfer...")
        const response = await this.oldHttpClient.request({url, method: "GET"})
        return response.data
    }
}