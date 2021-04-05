import com.google.gson.Gson
import org.java_websocket.client.WebSocketClient
import org.java_websocket.handshake.ServerHandshake

import java.io.PrintWriter
import java.net.{SocketException, URI}
import java.util.Date

class WebSocket {
  val webSocketClient: WebSocketClient = new WebSocketClient(new URI("ws://89.109.190.198:8003")) {
    override def onOpen(handshakedata: ServerHandshake): Unit = {
      send(new Gson().toJson(AuthWS("auth_req", "root", "123")))
    }

    override def onMessage(message: String): Unit = {
      if(message.contains("get_data_resp")) {
        println(message)
        new PrintWriter("kek") { write(message); close() }
      }
    }

    override def onClose(code: Int, reason: String, remote: Boolean): Unit = {
      throw new SocketException("socket close")
    }

    override def onError(ex: Exception): Unit = {
      ex.printStackTrace()
    }
  }
}


object WebSocket extends App{
  val kek = new WebSocket
  kek.webSocketClient.connectBlocking()
  Thread.sleep(1500)
  Seq("FFFFFF1000015F07").map(devEui => {
    val sendReq = new Gson().toJson(GetDataWS("get_data_req", devEui, Select((1616119200).toString, (new Date().getTime).toString, Int.MaxValue)))
    println(sendReq)
    kek.webSocketClient.send(sendReq)
  }
  )
  while (kek.webSocketClient.isOpen){

  }
}

case class AuthWS(cmd: String, login: String, password: String)
case class GetDataWS(cmd: String, devEui: String, select: Select)
case class Select(date_from: String, date_to: String, limit: Int)