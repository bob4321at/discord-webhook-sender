let send_button = document.getElementById("send");
let add_field_button = document.getElementById("add Field");
let field_list = document.getElementById("field_list")

let index_of_field = []
let value_fields  = []
let name_fields = []

add_field_button.onclick = function () {
  let name_input_field = document.createElement("input")
  let input_field = document.createElement("input")
  let index_field = document.createElement("h1")
  index_field.innerHTML = "Field " + (value_fields.length + 1)
  index_of_field.push(index_field)
  name_fields.push(name_input_field)
  value_fields.push(input_field)

  field_list.innerHTML = ""

  for (let i = 0; i < value_fields.length; i++) {
    field_list.appendChild(index_of_field[i])
    field_list.appendChild(name_fields[i])
    field_list.appendChild(value_fields[i])
  }
}

send_button.onclick = function () {
  let name = document.getElementById("get_name").value;
  messege = document.getElementById("messege").value;
  let footer_messege = document.getElementById("footer_input").value;
  let list_of_fields = []
  let fieldss = []

  for (let i = 0; i < value_fields.length; i++) {
    fieldss.push(name_fields[i].value)
    fieldss.push(value_fields[i].value)
    list_of_fields.push(fieldss)
    fieldss = []
  }

  fetch("/sendMessege", {
    "method": "POST",
    "body": JSON.stringify({"Url": document.getElementById("get_url").value, "Img_Url": document.getElementById("get_image").value, "Name": name, "Messegee": messege, "Fields": list_of_fields, "Footer": footer_messege})
  })
}
