const _asyncForEach = async (array, callback) => {
    for(let i = 0; i < array.length; i++){
        await callback(array[i])
    }
}

module.exports = {
    asyncForEach: _asyncForEach,
}