$(document).ready(function() {
	var list = (document).createElement('ul');
	list.id = "manager";
	addTaskToList = function(taskDescription){
		var list = document.getElementById("manager");
		var newElem = document.createElement('li');
		var newElemText = document.createElement('span');
		
		newElemText.innerHTML = taskDescription;
		newElem.appendChild(newElemText);
		list.appendChild(newElem);

		var delButton = document.createElement('button');
		delButton.innerHTML = "Удалить"
			
		delButton.onclick = function(){ list.removeChild(newElem); }

		newElem.appendChild(delButton);
	}
	
	document.getElementById("root").appendChild(list);
	addTaskToList("Сделать задание #3 по web-программированию");

	var addTaskInput = document.createElement('input');
	addTaskInput.type = "text"
	addTaskInput.id = "add_task_input"
	document.body.appendChild(addTaskInput);
	var addTaskButton = document.createElement("button");
	addTaskButton.id = "add_task"
	addTaskButton.innerHTML = 'Add task.'

	addTaskButton.onclick = function(){ addTaskToList(addTaskInput.value); }
	document.body.appendChild(addTaskButton);
});
