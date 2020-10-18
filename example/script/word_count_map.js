var isspace = c => c == " "

var MR_map = line => {
    let n = line.length
    for (let i = 0; i < n; ) { // Skip past leading whitespace 
        while ((i < n) && isspace(line[i])) i++;
        // Find word end 
        let start = i; 
        while ((i < n) && !isspace(line[i])) i++;
        if (start < i) {
            MR_mapEmit(line.substr(start,i-start),"1");
        }
    }
} 
