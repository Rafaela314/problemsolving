package main

/*
func (d *company) SearchCompanies(params models.Company) ([]models.Company, error) {

	fmt.Printf("\n MODEL CIA %v\n", params)

	var result []models.Company

	col, err := d.GetCollection()
	if err != nil {
		return nil, xerrors.Errorf("getting mongo connection %w", err)
	}

	var matchParams map[string]interface{} = make(map[string]interface{})

	var aggre []bson.M

	v := reflect.ValueOf(params)
	typeOfS := v.Type()
	st := reflect.TypeOf(params)

	for i := 0; i < v.NumField(); i++ {

		//Check element type

		if typeOfS.Field(i).Type.Kind() == reflect.Struct {
			//time will be here since it is a struct, as well as bson
			//	fmt.Printf("\n  É STRUCT ||  FIELD NAME %v\n", typeOfS.Field(i).Name)
			//	fmt.Printf("\n  TYYYPEEE %v\n", typeOfS.Field(i).Type.Kind())
		} else if typeOfS.Field(i).Type.Kind() == reflect.Slice || typeOfS.Field(i).Type.Kind() == reflect.Array {

			//	fmt.Printf("\n  É SLICE||  FIELD NAME %v\n", typeOfS.Field(i).Name)
			//fmt.Printf("\n  TYYYPEEE %v\n", typeOfS.Field(i).Type.Kind())

			if v.Field(i).Len() > 0 {
				var idList []interface{}

				//fmt.Printf("\n   É ARRAYYYYY FIELD NAME %v\n", typeOfS.Field(i).Name)
				//fmt.Printf("\n  LEN DO ARRAY %v\n", v.Field(i).Len())

				tag := strings.Split(st.Field(i).Tag.Get("bson"), ",")[0]
				//fmt.Printf("\n  FIELD NAME %v\n", typeOfS.Field(i).Name)

				matchParams[fmt.Sprintf("%s", tag)] = bson.M{"$in": idList}

				for i := 0; i < v.Field(i).Len(); i++ {
					idList = append(idList, v.Field(i).Interface())
				}
			}

		} else if typeOfS.Field(i).Type.Kind() == reflect.String {
			fmt.Printf("\n  É STRING||  FIELD NAME %v\n", typeOfS.Field(i).Name)

			/*
							 if id != "" && bson.IsObjectIdHex(id){
				        bsonObjectId := bson.ObjectIdHex(id)
				        return &bsonObjectId
					}*/ /*
			if typeOfS.Field(i).Name == "ID" {
				//if bson.IsObjectIdHex(fmt.Sprintf("%v", v.Field(i).Interface())) {
				fmt.Printf("\n  ÉBSON %v\n", v.Field(i).Interface())

			} else {
				if v.Field(i).Interface() != "" && v.Field(i).Interface() != "ObjectIdHex(\"\")" {
					fmt.Printf("\n  É STRING NAO VAZIA VAÇÇ %v\n", v.Field(i).Interface())
					fmt.Printf("\n  TYYYPEEE %v\n", typeOfS.Field(i).Type.Kind())
					tag := strings.Split(st.Field(i).Tag.Get("bson"), ",")[0]
					matchParams[fmt.Sprintf("%s", tag)] = typeOfS.Field(i).Name

				}
			}

		} else {
			if v.Field(i).Interface() != false && v.Field(i).Interface() != 0 && v.Field(i).Interface() != 0.0 && v.Field(i).Interface() != nil {
				fmt.Printf("\n  SOBROU NAME %v\n", typeOfS.Field(i).Name)
				fmt.Printf("\n FIELD INTERFACE %v\n", v.Field(i).Interface())
				fmt.Printf("\n  TYYYPEEE %v\n", typeOfS.Field(i).Type.Kind())
				fmt.Printf("\n  É STRING< NUMBER OU BOOL || FIELD NAME %v\n", typeOfS.Field(i).Name)

				tag := strings.Split(st.Field(i).Tag.Get("bson"), ",")[0]
				matchParams[fmt.Sprintf("%s", tag)] = typeOfS.Field(i).Name
			}

		}

	}


}*/
