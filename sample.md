```
import "math"

Complex : {
    float re, im

    new (float r, i) {
        re = r
        i = i
    }

    float magnitude() {
        return math.sqrt(re * re + im * im) 
    }

    float[] toArray() {
        return float[re, im]
    }
}
```