# What is the `Abstraction` ?
➜ Abstraction is used to add a layer of simplyfing to the clinet of our code <br>
➜ Abstraction is achieved by hiding complexities from the client and exposing essential methods only to the client <br>
➜ Abstraction is achieved by making the internal methods as private and exposing simple methods to be called by the client <br>

# What is the `Encapsulation` ?
➜ Encapsulation is concerned of : <br> <t>
- Data hiding <br>
- Separation of concern <br>
- Keeping the internal state of the object protected from the direct manipulation of the client <br>

➜ Example :  <br> <t>
- if we are building a bank account management system. <br>
- we shouldn't let the balance to be public because user can then say `bank.balance = -100`. <br>
- so we shuold set the balance as `private` and expose from this class 2 methods one for deposite and one for withdraw and within these two methods we should handle all the comlexities and possible conditions to deposite or withdraw money. <br>
- then we should implement `getter` for the client to be able to test against our balance value. <br>
- encapsulation provides for us the `data integrity and validity`. <br>

# What is the `Polymorphism` ?
➜ Polymorphism : Many-forms .<br>
➜ Its simply the ability of one method to work with different data types. <br>
➜ Its a `superclass` has `many forms`. <br>
➜ we have 3-types of polymorphism :  <br> <t>
- `subtype polymorhpism` : inheritence <br>
- `parametric polymorphism` : generic <br>
- `Ad hoc polymorphism` : function/operator overloading <br>

➜ polymorphism is the ability of instances of subclasses to be treated as instances of superclass <br>

```ts
interface Shape {
    area() : number
    perimeter() : number
}

class Circle implements Shape {
    constructor(private redius : number ){
        this.redius = redius
    }


    area() : number {
        return this.redius * Math.PI * Math.PI
    }

    perimeter() : number {
        return 2 * this.redius * Math.PI 
    }
}



class Rectangle implements Shape {
    constructor(private w : number, private l : number ){
        this.w = w
        this.l = l
    }


    area() : number {
        return this.w * this.l
    }

    perimeter() : number {
        return 2 * this.w * this.l
    }
}


// the method accepts instance of superclass
function CalculateShapeArea( shape : Shape ) : number {
    return shape.area()
}

const circle : Circle = new Circle(5)
// but we use it with instances of different subclasses implements this superclass
console.log((CalculateShapeArea(circle)))

const rect : Rectangle = new Rectangle(5, 4)
// but we use it with instances of different subclasses implements this superclass
console.log(CalculateShapeArea(rect));
```

➜ Polymorphism is achieved in express .. ! <br>
in the following example you will see that you can define different middleware as long as each middleware function takes the 3 params defined in the middleware interface
```ts
const app = express() 

const middleware_1 = (req:Request, res:Response, next:NextFunction) => {
    // do some custom work
    next()
}

const middleware_2 = (req:Request, res:Response, next:NextFunction) => {
    // do some custom work
    next()
}

app.use(middleware_1)
app.use(middleware_2)
```

# What is the `SRP` ?
➜ SRP stands for Single Responsibility Pattern <br>
➜ A class should only have one reason to change <br>
`User` class should only be concerned of managing the attributes of the user, but we shouldn't define any method related to authenticating the user.
```ts
class User {
    constructor(
        private name : string, 
        private age : number
    ){}

    // getters and setters to deal with user properties
}

class UserAuth {
    constructor(
        private u : User
    ){}

    register() : bool {
        // impl of registeration of an user
    }
}
```

# What is the `OCP` ?
➜ OCP stands for Open Closed Principl, which means the functionality is open for extension closed for modification <br>
➜ Bad Design : <br>
```ts
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
```
now if we need to add new type of users, we will open the `giveDiscount` function and modify it which is wrong. <br>

➜ Good Design : <br>
```ts
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
```