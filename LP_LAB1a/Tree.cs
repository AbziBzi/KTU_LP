namespace LP_LAB1a
{
    public class Tree
    {
        public string Type { get; private set; }
        public int Age { get; private set; }
        public double Height { get; private set; }
        public long result;

        public Tree(string type, int age, double height)
        {
            this.Type = type;
            this.Age = age;
            this.Height = height;
        }

        private long FindPrimeNumber()
        {
            int n = (int)(this.Height * this.Age * this.Type.Length);
            int count = 0;
            long a = 2;
            while (count < n)
            {
                long b = 2;
                int prime = 1;// to check if found a prime
                while (b * b <= a)
                {
                    if (a % b == 0)
                    {
                        prime = 0;
                        break;
                    }
                    b++;
                }
                if (prime > 0)
                {
                    count++;
                }
                a++;
            }
            this.result = (--a);
            return (--a);
        }

        public int CompareTo(Tree tree)
        {
            if (tree.FindPrimeNumber() > this.FindPrimeNumber())
                return -1;
            if (tree.FindPrimeNumber() < this.FindPrimeNumber())
                return 1;
            return 0;
        }

        public long ReturnResult()
        {
            return result;
        }
    }
}