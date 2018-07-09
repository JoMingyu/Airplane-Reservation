using System;
using System.Collections;

namespace AIRPLAINE_RESERVATION
{
    public class Airplane
    {
        private ArrayList departmentTimesInHour;
        private int seatCount;

        private string airline;

        private readonly ArrayList reservations;
        private readonly ArrayList reservationWaitings;
        
        public Airplane(ArrayList departmentTimesInHour, int seatCount, string airline)
        {
            this.departmentTimesInHour = departmentTimesInHour;
            this.seatCount = seatCount;

            if(Airline.airlines.Contains(airline))
            {
                this.airline = airline;
            }
            else
            {
                throw new ArgumentException("What's this airline?");
            }
        }

        public bool AddNewReservation(Person p)
        {
            /* 새로운 사람을 예약 */

            bool CanReservation()
            {
                /* 예약 가능한 상태인지 확인 */

                if (this.reservations.Count < this.seatCount)
                {
                    // 예약자가 전체 자리수보다 적은 경우
                    return true;
                }
                else
                {
                    return false;
                }
            }

            if(CanReservation())
            {
                this.reservations.Add(p);

                return true;
            }
            else
            {
                // 이미 모두 예약되어 대기자에 추가
                this.reservationWaitings.Add(p);

                return false;
            }
        }

        public bool CancelReservation(Person p)
        {
            /* 예약 철회 */

            bool CanCancel()
            {
                /* 철회가 가능한 상태인지(예약을 하긴 했는지) 확인 */

                if(this.reservations.Contains(p))
                {
                    // 예약자 목록에 이놈이 존재하는 경우
                    return true;
                }
                else
                {
                    return false;
                }
            }

            if(CanCancel())
            {
                this.reservations.Remove(p);

                var res = this.AddNewReservation((Person)this.reservationWaitings[0]);
                // 대기자 추가

                if (res)
                {
                    // 예약에 성공하면

                    this.reservationWaitings.RemoveAt(0);
                    // 이 친구를 대기자에서 제거
                }

                return true;
            }
            else
            {
                return false;
            }
        }

        public int InquireReservation(Person p)
        {
            if(this.reservations.Contains(p))
            {
                return this.reservations.IndexOf(p);
                // 예약이 돼있으면 그 자리의 인덱스 반환
            }
            else
            {
                return 0;
                // 안돼있으면 0
            }
        }
    }
}
