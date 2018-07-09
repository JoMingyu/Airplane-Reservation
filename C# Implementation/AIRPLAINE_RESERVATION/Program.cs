using System;
using System.Collections;

namespace AIRPLAINE_RESERVATION
{
    class MainClass
    {
        public static void Main(string[] args)
        {
            Airline.airlines.Add("jeju");
            Airline.airlines.Add("new york");
            Airline.airlines.Add("sidney");

            ArrayList jejuAirs = new ArrayList(
                new Airplane[] {
                    new Airplane(new ArrayList(new int[] { 9 }), 50, "jeju"),
                    new Airplane(new ArrayList(new int[] { 13 }), 50, "jeju"),
                    new Airplane(new ArrayList(new int[] { 22 }), 50, "jeju")
                }
            );

            ArrayList newYorkAirs = new ArrayList(
                new Airplane[] {
                    new Airplane(new ArrayList(new int[] { 9 }), 200, "new york"),
                    new Airplane(new ArrayList(new int[] { 13 }), 200, "new york"),
                    new Airplane(new ArrayList(new int[] { 22 }), 200, "new york")
                }
            );

            ArrayList sidneyAirs = new ArrayList(
                new Airplane[] {
                    new Airplane(new ArrayList(new int[] { 9 }), 200, "sidney"),
                    new Airplane(new ArrayList(new int[] { 13 }), 200, "sidney"),
                    new Airplane(new ArrayList(new int[] { 22 }), 200, "sidney")
                }
            );

            Console.WriteLine("");

            int choice = int.Parse(Console.ReadLine());

            switch (choice)
            {
                default:
                    break;
            }
        }
    }
}
