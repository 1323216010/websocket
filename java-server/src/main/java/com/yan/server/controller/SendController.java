package com.yan.server.controller;

import com.yan.server.Endpoint.WsServerEndpoint;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;
import java.util.Map;

/**
 * @author yan
 **/
@RestController
public class SendController {
    @GetMapping("/send")
    public void send(@RequestParam Map<String, Object> map) throws IOException {
        WsServerEndpoint.wsCoon.get(Integer.valueOf(String.valueOf(map.get("id")))).getBasicRemote().sendText(String.valueOf(map.get("msg")));
    }
}
