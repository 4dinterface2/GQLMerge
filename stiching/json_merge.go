package stiching

type ResultJson struct {
	result interface{}
}

func (j *ResultJson) AddMap(source interface{}){
	j.result = merge( &j.result, &source)
	//js, _:=json.Marshal(&j.result)
	//fmt.Println("res = ",  string(js) )
}

//object and array merge
func merge(dst *interface{}, src *interface{}) *interface{}{

	switch target := (*dst).(type) {
		case map[string]*interface{}:
			for key, srcVal := range (*src).(map[string]interface{}) { //TODO не факт что в src map
				if dstVal, ok:= target[key]; ok {
					target[key] = merge(dstVal, &srcVal )
				} else {
					target[key] = &srcVal //копируем ссылку
				}
			}
			return  dst
			break

		case []interface{}:
			for _, srcVal := range (*src).([]interface{}) { //TODO не факт что в src map
				target = append(target, srcVal) //в самом простом случае элементы просто добавляются в массив
				//более сложные могут включать в себя поиск элементов по ключу
			}
			return dst
			break

		default:
			return src
			break
	}
	return nil
}