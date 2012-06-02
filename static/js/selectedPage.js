$(document).ready(function(){
	$(function(){
   	var path = location.pathname.substring(1);
   	if ( path )
     		$('nav a[href$="' + path + '"]').attr('class', 'selected');
	});
});