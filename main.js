const fs = require('fs')
const readline = require('readline');

var input = fs.readFileSync('./day5.txt').toString().split('\n')
var lines = {}
var start, end

function coeffDir(x1, x2, y1, y2) {
    coeff= (y1-y2)/(x1-x2)
    if (coeff===1 || coeff===-1){
        return coeff
    }
    return 0
}

input.forEach(element => {
    coords = element.split(/(?:,| -> )+/)
    x1 = parseInt(coords[0])
    x2 = parseInt(coords[2])
    y1 = parseInt(coords[1])
    y2 = parseInt(coords[3])

    if (x1 === x2) {
        if (y1 <= y2) {
            start = y1
            end = y2
        } else {
            end = y1
            start = y2
        }
        for (let i = start; i <= end; i++) {
            if (!lines[x1+","+i]) {
                lines[x1+","+i] = 1
            } else {
                lines[x1+","+i] = lines[x1+","+i]+1
            }
        } 
    } else if (y1 === y2) {
        if (x1 <= x2) {
            start = x1
            end = x2
        } else {
            end = x1
            start = x2
        }
        for (let i = start; i <= end; i++) {
            if (!lines[i+","+y1]) {
                lines[i+","+y1] = 1
            } else {
                lines[i+","+y1] = lines[i+","+y1] + 1
            }
        } 
    } else if (coeffDir(x1, x2, y1, y2) !== 0) {
        starty=0
        if (x1 <= x2) {
            start = x1
            starty = y1
            end = x2
        } else {
            end = x1
            start = x2
            starty = y2
        }
        if (coeffDir(x1, x2, y1, y2) ===1) {
            i = 0
            while (start+i <= end) {
                if (!lines[(start+i)+","+(starty+i)]) {
                    lines[(start+i)+","+(starty+i)] = 1
                } else {
                    lines[(start+i)+","+(starty+i)] = lines[(start+i)+","+(starty+i)] + 1
                }
                i++
            } 
        } else {
            i = 0
            while (start+i <= end) {
                if (!lines[(start+i)+","+(starty-i)]) {
                    lines[(start+i)+","+(starty-i)] = 1
                } else {
                    lines[(start+i)+","+(starty-i)] = lines[(start+i)+","+(starty-i)] + 1
                }
                i++
            } 
        }

    }
});

count = 0

for (const [key, val] of Object.entries(lines)) {
    if (val > 1) {
        count++
    }
}

console.log(count)
