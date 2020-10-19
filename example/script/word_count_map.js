var isspace = function(c) { return c == " " }

var MR_map = function(line) {
    var n = line.length
    for (var i = 0; i < n; ) { // Skip past leading whitespace 
        while ((i < n) && isspace(line[i])) i++;
        // Find word end 
        var start = i; 
        while ((i < n) && !isspace(line[i])) i++;
        if (start < i) {
            MR_mapEmit(line.substr(start,i-start),"1");
        }
    }
} 
