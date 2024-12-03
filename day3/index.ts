import * as fs from "fs";
import * as path from "path";

const filePath = path.join(__dirname, "input.txt");

try {
    const data = fs.readFileSync(filePath, 'utf-8')
    const tokens = data.match(/mul\(\d+,\d+\)|do\(\)|don't\(\)/g)
    let multiply = true;
    let sum = 0;
    console.log(tokens)
    tokens?.map((token) => {
        if (token === "do()") {
            multiply = true;
        } else if (token === "don't()") {
            multiply = false;
        } else if (token.startsWith("mul(")) {
            const [_, a, b] = token.match(/mul\((\d+),(\d+)\)/) || [];
            if (multiply) {
                sum += Number(a) * Number(b); 
            }
        } 
    })  
    console.log(sum);
}  catch (error) {
       console.error(error)
}