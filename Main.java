import com.sun.net.httpserver.*; import java.net.*; import java.io.*; import java.util.*;
public class Main { public static void main(String[] a) throws Exception {
  System.out.println("STARTUP-MARKER java");
  HttpServer s=HttpServer.create(new InetSocketAddress("0.0.0.0",8080),0);
  s.createContext("/",x->{ if(x.getRequestURI().getPath().equals("/log")) System.out.println("E2E-LOG-MARKER");
    StringBuilder e=new StringBuilder("{");
    for(Map.Entry<String,String> k:System.getenv().entrySet()) if(k.getKey().startsWith("E2E_")) e.append("\"").append(k.getKey()).append("\":\"").append(k.getValue()).append("\",");
    String b="{\"lang\":\"java\",\"env\":"+(e.length()>1?e.substring(0,e.length()-1):e)+"}}";
    x.getResponseHeaders().add("content-type","application/json"); byte[] o=b.getBytes(); x.sendResponseHeaders(200,o.length); x.getResponseBody().write(o); x.close(); });
  s.start(); } }
