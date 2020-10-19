
var MR_reduce = function(values) {
    var value = 0
    values.forEach(function(v) {
        value += parseInt(v)
    });
    // Emit sum for input->key() 
    MR_reduceEmit(value);
}
