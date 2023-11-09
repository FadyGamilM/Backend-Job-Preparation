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


function CalculateShapeArea( shape : Shape ) : number {
    return shape.area()
}

const circle : Circle = new Circle(5)
console.log((CalculateShapeArea(circle)))

const rect : Rectangle = new Rectangle(5, 4)
console.log(CalculateShapeArea(rect));



