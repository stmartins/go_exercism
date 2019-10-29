package flatten

func Flatten(input interface{}) []interface{} {

    result := make([]interface{}, 0);

    for _, val := range input.([]interface{}) {
        switch val.(type) {
            case []interface{}:
                result = append(result, Flatten(val)...)
            case int:
                result = append(result, val)
        }
    }
    return result;
}
