import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class finalprac {
    public static void main(String[] args) {
        List<filter>list=new ArrayList<>();

        filter fi=new filter(12, "amu", "kathamndu","ssshar");
        filter fi1=new filter(12, "amu", "banepa","askjka");
        filter fi2=new filter(12, "amu", "banepa","oqwjeowe");

        list.add(fi);

        list.add(fi1);
        list.add(fi2);

        List<mapping>lists=list.stream().filter(a->a.address.equalsIgnoreCase("kathamndu")).map(a->{
mapping ma=new mapping();

ma.setAddress(a.address);
ma.setName(a.name);

return ma;
        }).collect(Collectors.toList());

        for(mapping li:lists){
            System.out.println("first"+li.address);
            System.out.println("third"+li.name);
        }




    }
    
}
