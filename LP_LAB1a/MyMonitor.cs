using System;
using System.Collections.Generic;
using System.Threading;

namespace LP_LAB1a
{
    public class MyMonitor
    {
        private static readonly Object myLock = new Object();
        private Tree[] data;
        private int count;
        private bool stillWorking = true;

        public MyMonitor(int size)
        {
            this.data = new Tree[size];
        }

        public void AddItem(Tree item)
        {
            lock (myLock)
            {
                while (this.count == data.Length)
                {
                    Monitor.Wait(myLock);
                }
                int index = FindIndex(item);
                if (data[index] != null)
                {
                    ShiftElement(index);
                }
                data[index] = item;
                count++;
                Monitor.PulseAll(myLock);
            }
        }

        private int FindIndex(Tree item)
        {
            int index = 0;

            for (int i = 0; i < data.Length; i++)
            {
                if (data[i] == null)
                    break;
                if (data[i].CompareTo(item) <= 0)
                    index = i + 1;
                else
                    break;
            }
            return index;
        }

        private void ShiftElement(int index)
        {
            for (int i = data.Length - 1; i > index; i--)
            {
                data[i] = data[i - 1];
            }
            data[index] = null;
        }

        public Tree Pop()
        {
            lock (myLock)
            {
                while (count == 0)
                {
                    if (stillWorking)
                        Monitor.Wait(myLock);

                    else if (count == 0)
                    {
                        Monitor.PulseAll(myLock);
                        return null;
                    }

                }
                Tree item = data[count - 1];
                data[count - 1] = null;
                count--;
                Monitor.PulseAll(myLock);
                return item;
            }
        }

        public Tree GetTree(int index)
        {
            lock (myLock)
            {
                if (index < 0 || index > data.Length)
                    throw new ArgumentException("invalid index in Monitor.GetTree(int index)");
                return data[index];
            }
        }

        public int GetCount()
        {
            return this.count;
        }

        public void SetStillWorking(bool isWorking)
        {
            this.stillWorking = isWorking;
        }

        public bool GetStillWorking()
        {
            return stillWorking;
        }
    }
}