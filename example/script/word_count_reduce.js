
var MR_reduce = values => {
    let value = 0
    values.forEach(v => {
        value += parseInt(v)
    });
    // Emit sum for input->key() 
    MR_reduceEmit(value);
}
