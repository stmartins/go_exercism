package accumulate

func Accumulate(input []string, fctn func(string) string) []string {

    var ret = []string{}

    for _, v := range input {
        ret = append(ret, fctn(v))
    }
    return ret;
}
