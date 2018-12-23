$(".tag").on('click', function (event) {
    id = event.target.id;
    if(id == "tag_X") {
        $(".tile.is-child:has(.tag)").css("opacity", 1)
        $(".tag_X").hide()
        return
    }
    $(".tile.is-child:has(.tag)").css("opacity", 0.5)
    $(".tile.is-child:has(.tag." + id + ")").css("opacity", 1)
    $(".tag_X").show()
});