$(document).ready(function(){
    $(".cell").on("click", function(){
        if($(this).hasClass("active")) {
            $(this).find("input[type=hidden]").val("false")
        } else {
            $(this).find("input[type=hidden]").val("true")
        }
        $(this).toggleClass("active")
        var colIndex = $(this).prevAll().length;
        var rowIndex = $(this).parent('.matrix-row').prevAll().length;

        $.ajax({
            type: "post",
            url: "/move",
            data: {
                row: rowIndex,
                col: colIndex,
                hash: sourceHash
            },
            dataType: 'json',
            success: function(response) {
                if(!response.success) {
                    alert(response.error)
                } else {
                    for(var i = 0; i < response.board.length; i++) {
                        for(var j = 0; j < response.board[i].length; j++) {
                            if(response.board[i][j]) {
                                $(".grid-board .matrix-row:nth-child(" + i + ")").find(".cell:nth-child(" + j + ")").addClass("active");
                            } else {
                                $(".grid-board .matrix-row:nth-child(" + i + ")").find(".cell:nth-child(" + j + ")").removeClass("active");
                            }
                        }
                        console.log(response.board[i])
                        console.log("---------------")
                    }
                }
            }
        })
    });
});