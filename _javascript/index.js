function tagClick(e) {
    id = $(this).attr("id");
    if(id === "tag_X") {
        $(".tile.is-child:has(.tag)").css("opacity", 1)
        // $(".tag_X").hide()
        return
    }
    $(".tile.is-child:has(.tag)").css("opacity", 0.2)
    $(".tile.is-child:has(.tag." + id + ")").css("opacity", 1)
    // $(".tag_X").show()
}

$(".tag_X").on('click', tagClick);
$(".tag").on('click', tagClick);


$(".tile.is-child").on('click', function(e) {
    if ($(this).css("opacity") !== "1") {
        return
    }
    if ($(e.target).hasClass("tag")) {
        return
    }

    if ($(e.target).closest(".modal").length > 0) {
        return
    }

    m = $(this).find(".modal")
    m.addClass("is-active")
})

$(".modal-background").on('click', function(e){
    $(this).parent().removeClass("is-active")
})

$(".modal-delete").on('click', function(e){
    $(this).closest(".modal").removeClass("is-active")
})