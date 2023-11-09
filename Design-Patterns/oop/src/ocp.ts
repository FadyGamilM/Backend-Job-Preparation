 // open closed principle 

 //! [X] Bad Design 
class User {
    constructor(private username: string, private _role : string ){}

    // HINT : the field must be named differently than the getter property ...
    public get role () : string {
        return this._role
    } 
}

class Discount {
    giveDiscount(user: User) : number{
        if (user.role == "ADMIN") {
            return 50
        }else if (user.role == "REGULAR"){
            return 20
        }else {
            return 10
        }
    }
}


// Good Design 
interface IUser {
    username : string 
    role : string 

    getDiscount() : number
}

class RegularUser implements IUser {
    username: string; 
    role: string;
    constructor(username :string ){
        this.role = "REGULAR"
        this.username = username;
    }

    // the impl of the discount is set here not conditonially in the DIscount class 
    public getDiscount() : number { 
        return 20
    }
}


class AdminUser implements IUser {
    username: string; 
    role: string;
    constructor(username :string ){
        this.role = "ADMIN"
        this.username = username;
    }

    // the impl of the discount is set here not conditonially in the DIscount class 
    public getDiscount() : number { 
        return 50
    }
}

class Discount_V2 {
    giveDiscount(user : IUser) {
        return user.getDiscount()
    }
}


