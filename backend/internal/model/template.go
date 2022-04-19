package model

type Template struct {
	ID     int
	Name   string
	Fields map[string]interface{}
	// range th/td tags
	TempId int
}

type Templ1 struct {
	ID           int // tmpl id
	Title        string
	Description  string
	Company      // get list
	Departament  // get list
	Status       // get list
	CreationDate string
	Author       User `json:"auhtor_id"`  // get jwt token
	ToWhom       User `json:"to_whom_id"` // return list user
	// TempalteId   int  `json:"template_id"` //
}

// create sample 1,2,3,4,5 html
// create struct 1,2,3,4,5

// 1 create -> return  list TempalteName
// 2 client choice  template X -> return structX json
// td, th, generate reflect -> map[string]string

/*
in further create template : admin -> create template
1.1 create temp:  Name tempalte
1.2 create fields:  key:title:type:tempId: tag
1.3 generate by Fields - html -> <h1> {{ .Key }} </h1>
1.4 return client from Db by templID -> json {title: key}
-----------------

1 client:  /create pdf doc -> getFields By TempId -> map[jsonKey]title
2 backend: parseHtml(tmlpId.html, map[jsonKey]title)
3 save docTable - path, authorId
4 save fileServer -> file
5 back -> return client -> list Users
6 return list  client - statusses
6.1 sign - set statusId(default 1 (pending)) - toWhomId, docId - pending

7 add events history table
-------------------

MyDocs -> byIdDoc -> history

AtSign - sent/incoming/done

1 AtSign -> sent -> show list docs - status pending

2 AtSign -> incoming - list me docs - status pending

3 AtSign -> incoming - getDocById - sign - ecp -> change status - success
3.1  move to sent;
3.2 add  history  event

2
*/
