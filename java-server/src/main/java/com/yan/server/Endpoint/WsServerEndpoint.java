package com.yan.server.Endpoint;

import jakarta.websocket.OnClose;
import jakarta.websocket.OnMessage;
import jakarta.websocket.OnOpen;
import jakarta.websocket.Session;
import jakarta.websocket.server.ServerEndpoint;
import org.springframework.stereotype.Component;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * @author yan
 **/
@ServerEndpoint("/")
@Component
public class WsServerEndpoint {
    private Session session;

    public static List<Session> wsCoon = new ArrayList<Session>();
    /**
     * 连接成功
     *
     * @param session
     */
    @OnOpen
    public void onOpen(Session session) {
        this.session = session;
        wsCoon.add(session);
    }

    /**
     * 连接关闭
     *
     * @param session
     */
    @OnClose
    public void onClose(Session session) {
        System.out.println("连接关闭");
    }

    /**
     * 接收到消息
     *
     * @param str
     */
    @OnMessage
    public String onMsg(String str) throws IOException {
        return str;
    }

}