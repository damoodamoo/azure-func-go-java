package com.damoo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;

import com.damoo.models.*;
import com.damoo.repos.ShoppingList;

@RestController
public class ShoppingAPI {

  @Autowired
  private ShoppingList repo;

  @PostMapping(path = "/add")
  public void addItem(@RequestBody ShoppingItem item) {
    System.out.println("JAVA: Adding item");
    repo.Add(item);
  }

  @RequestMapping("/get")
  public ShoppingItem getItem(@RequestParam("id") int id) {
    System.out.println("JAVA: Getting item");

    return repo.Get(id);
  }

  @RequestMapping("/list")
  public ArrayList<ShoppingItem> getList() {
    System.out.println("JAVA: Getting list");
    return repo.List();
  }

  @RequestMapping("/send-items")
  public InvokeResponse sendItems() {
    System.out.println("JAVA: Sending items... somewhere");

    InvokeResponse resp = new InvokeResponse();
    resp.Outputs.put("output1", repo.List());
    return resp;
  }

  @PostMapping("/process-items")
  public InvokeResponse processItems(@RequestBody InvokeRequest req) {

    System.out.println("JAVA: Received items and sending... elsewhere...");

    // we get a list of items, and do *something*
    // ... then push them somewhere else...
    
    InvokeResponse resp = new InvokeResponse();
    resp.Outputs.put("output1", req.Data);
    resp.Outputs.put("output2", req.Data);
    return resp;
  }
}