using System;
namespace AIRPLAINE_RESERVATION
{
    public class Person
    {
        public string Id { get; }
        public string Pw { get; set; }
        public string Name { get; set; }

        public Person(string id, string pw, string name)
        {
            this.Id = id;
            this.Pw = pw;
            this.Name = name;
        }
    }
}
