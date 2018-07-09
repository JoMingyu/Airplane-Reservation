using System;
using System.Collections;
using System.Collections.Generic;

namespace AIRPLAINE_RESERVATION
{
    public static class ReservationSystem
    {
        private static Dictionary<string, Person> users;

        public static bool Signup(string id, string pw, string name)
        {
            bool IsIdExist()
            {
                if(users.ContainsKey(id))
                {
                    return true;
                }
                else
                {
                    return false;
                }
            }

            if(IsIdExist())
            {
                return false;
            }
            else
            {
                users.Add(id, new Person(id, pw, name));

                return true;
            }
        }

        public static Person auth(string id, string pw)
        {
            if(users.ContainsKey(id))
            {
                Person user = users[id];

                if(user.Pw == pw)
                {
                    return user;
                }
            }

            return null;
        }
    }
}
