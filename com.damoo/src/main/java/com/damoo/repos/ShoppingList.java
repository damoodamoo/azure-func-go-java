package com.damoo.repos;

import java.util.ArrayList;
import com.damoo.models.ShoppingItem;
import org.springframework.stereotype.Repository;

@Repository
public class ShoppingList {

    private ArrayList<ShoppingItem> Items;

    public ShoppingList(){
        Items = new ArrayList<ShoppingItem>();
    }

    public void Add(ShoppingItem item){
        Items.add(item);
    }

    public ShoppingItem Get(int id){
        return Items.get(id);
    }

    public ArrayList<ShoppingItem> List(){
        return Items;
    }

}