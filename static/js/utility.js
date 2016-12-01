/**********************************
 * Project:  blob_storage
 * Author:   Ray Chen
 * Email:    raychen0411@gmail.com
 * Time:     11-29-2016
 * All rights reserved!
 ***********************************/

"use strict";
function create_slider_for_tc() {
    $( "#pkg_loss_slider" ).slider({
        range: "max",
        min: 0,
        max: 100,
        value: 0,
        slide: function( event, ui ) {
            $( "#pkg_loss_amount" ).val( ui.value );
        }
    });

    $( "#pkg_latency_slider" ).slider({
        range: "max",
        min: 0,
        max: 100,
        value: 0,
        slide: function( event, ui ) {
            $( "#latency_amount" ).val( ui.value );
        }
    });

    $( "#pkg_reorder_slider" ).slider({
        range: "max",
        min: 0,
        max: 100,
        value: 0,
        slide: function( event, ui ) {
            $( "#reorder_amount" ).val( ui.value );
        }
    });
}

function ajaxCall(type, val) {
    $.ajax({
        url: "/update/tc",
        data: {
            type: type,
            value: val
        },
        success: function( result ) {
            console.log("finish Ajax call");
            $( "#system_output" ).html( result );
        }
    });
}

function update_pkg_loss() {
    var val = $("#pkg_loss_amount").val();
    ajaxCall("1", val);
}

function update_latency() {
    var val = $("#latency_amount").val();
    ajaxCall("2", val);
}

function update_reorder() {
    var val = $("#reorder_amount").val();
    ajaxCall("3", val);
}

function show_ping() {
    $.ajax({
        url: "/status/ping",
        data: {
        },
        success: function( result ) {
            console.log("finish Ajax call");
            $( "#system_output" ).html( result );
        }
    });
}

function show_cpu() {
    $.ajax({
        url: "/status/cpu",
        data: {
        },
        success: function( result ) {
            console.log("finish Ajax call");
            $( "#system_output" ).html( result );
        }
    });
}

function show_memory() {

}

function show_bandwidth() {

}