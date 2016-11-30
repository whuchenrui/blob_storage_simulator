/**********************************
 * Project:  blob_storage
 * Author:   Ray Chen
 * Email:    raychen0411@gmail.com
 * Time:     11-29-2016
 * All rights reserved!
 ***********************************/

"use strict";
function update_pkg_loss() {
    var val = $("#pkg_loss_amount").val();
    alert(val);
    $.ajax({
        url: "/update/tc",
        data: {
            type: "Ajax",
            value: val
        },
        success: function( result ) {
            alert("get result"+result);
            $( "#system_output" ).html( result );
        }
    });
}


function update_latency() {
    var val = $("#latency_amount").val();
    alert(val);
}

function show_ping() {

}

function show_cpu() {

}

function show_memory() {

}

function show_bandwidth() {

}