interface ShapeProps {
    color: string
    x : number
    y : number
}

abstract class Shape {
    constructor (
        public props : ShapeProps
    ){}

    abstract clone() : Shape
}

class Circle extends Shape{
    constructor(
        public radius : number,
        public props : ShapeProps
    ){
        super(props)
    }

    public clone() : Shape{
        return new Circle(this.radius, {
            color: this.props.color,
            x: this.props.x,
            y: this.props.y,
        });
    }
}

class Rectangle extends Shape{
    constructor(
        public width : number,
        public height : number,
        public props : ShapeProps
    ){
        super(props)
    }

    public clone() : Shape{
        return new Rectangle(this.width, this.height, {
            color: this.props.color,
            x: this.props.x,
            y: this.props.y,
        });
    }
}

// client code .. 