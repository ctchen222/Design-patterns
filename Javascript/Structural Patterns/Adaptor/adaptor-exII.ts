interface BackendUser {
    firstName: string
    lastName: string
    age: number
}

interface UserProfile {
    fullName: string
    age: number
}

export class UserAdaptor {
    static transform(user: BackendUser): UserProfile {
        return {
            fullName: `${user.firstName} ${user.lastName}`,
            age: user.age
        }
    }
}