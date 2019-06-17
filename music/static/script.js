async function send(name) {
	let response = await fetch('/api?q='+name);
	var Jresponce=JSON.parse(await response.text());
	var length = Object.keys(Jresponce).length; 
	console.log("len=",length);
	var i = 0;
	var send="";
	var play_button=`<button id="play">Прослушать </button>`;
	while (i < length) {
		var post = `<p class="box"> `+Jresponce[i]["artist"] +" - " +Jresponce[i]["title"]+`<a href="/api?link=`+"cut="+Jresponce[i]["owner_id"]+"cut="+Jresponce[i]["artist"]+"-"+Jresponce[i]["title"]+"cut="+Jresponce[i]["url"]+`"><span class="music_text">`+play_button+`</span></a> </p>"`;
		console.log(Jresponce[i]);
		send=send+post;
		i++;
		// send_2=send_2+`
	}
    document.querySelector('#lol').innerHTML =send;
	
}
function alarms(form){
    var name = form.searchString.value;
	send(name);
}     
