using System;
using System.Collections.Generic;
using System.IO;
using Newtonsoft.Json.Linq;
using System.Threading;
using System.Diagnostics;

namespace LP_LAB1a
{
    class Program
    {
        private static readonly int threadsCount = 25;
        // private static readonly string file = "./data/IFF72_ZubowiczE_L1_dat_1.json";
        // private static readonly string file = "./data/IFF72_ZubowiczE_L1_dat_2.json";
        private static readonly string file = "./data/IFF72_ZubowiczE_L1_dat_3.json";
        private static readonly string resultsPath = "./data/IFF72_ZubowiczE_L1_rez_1.txt";
        private static readonly int MONITOR_ELEMENTS_COUNT = 25;
        private static readonly long FILTER_VALUE = 92941;
        private static List<Thread> threads = new List<Thread>();
        private static List<Tree> primamyTreeList = new List<Tree>();
        private static MyMonitor workingTrees = new MyMonitor(MONITOR_ELEMENTS_COUNT);
        private static MyMonitor finalTrees = new MyMonitor(50);

        static void Main(string[] args)
        {
            ReadJsonFile(file, primamyTreeList);
            Stopwatch sw = new Stopwatch();
            sw.Start();

            // Creating threads
            for (int i = 0; i < threadsCount; i++)
            {
                threads.Add(new Thread(() => Execute()));
                threads[i].Start();
            }
            // Adding data to Monitor
            AddTreesToMonitor();
            //Joining all threads
            foreach (var thread in threads)
            {
                thread.Join();
            }

            sw.Stop();
            WriteResultsToFile();
            System.Console.WriteLine("Executed in: {0,-20} with {1} threads", sw.Elapsed, threadsCount);

        }

        /// Taking last element of Monitor and checking if it's result is less or equal 
        /// than FILTER_VALUE, if yes it adds to final data Monitor
        private static void Execute()
        {
            while (workingTrees.GetCount() != 0 || workingTrees.GetStillWorking())
            {
                Tree tree = workingTrees.Pop();

                if (tree != null)
                {
                    long res = tree.ReturnResult();
                    tree.result = res;
                    if (tree.result <= FILTER_VALUE)
                    {
                        finalTrees.AddItem(tree);
                    }
                }
            }
        }

        /// Adding Trees objects from List to Monitor.
        /// After adding data, setting stillWorking to False
        private static void AddTreesToMonitor()
        {
            foreach (var tree in primamyTreeList)
            {
                workingTrees.AddItem(tree);
            }
            workingTrees.SetStillWorking(false);
        }


        /// Reading from json FILE and adding to List<Tree> data struct
        /// string filePath - path of json file with data
        private static void ReadJsonFile(string filePath, List<Tree> trees)
        {
            using (StreamReader r = new StreamReader(filePath))
            {
                var data = r.ReadToEnd();
                JToken token = JObject.Parse(data);
                foreach (var item in token.SelectToken("trees"))
                {
                    string type = (string)item.SelectToken("type");
                    int age = (int)item.SelectToken("age");
                    double height = (double)item.SelectToken("height_m");
                    trees.Add(new Tree(type, age, height));
                }
            }
        }

        /// Writes primary data to File, and to the same file writes
        /// writes results that are filtered by result parameter
        private static void WriteResultsToFile()
        {
            if (File.Exists(resultsPath))
                File.Delete(resultsPath);

            using (StreamWriter writer = new StreamWriter(resultsPath))
            {
                writer.WriteLine("Pradiniai duomenys:");
                for (int i = 0; i < primamyTreeList.Count; i++)
                {
                    writer.WriteLine("| {0,-5}| {1,-15}| {2,-10}| {3,-10}|", i + 1, primamyTreeList[i].Type,
                                     primamyTreeList[i].Age, primamyTreeList[i].Height);
                }
                writer.WriteLine("\nRezultatai:");
                for (int i = 0; i < finalTrees.GetCount(); i++)
                {
                    var tree = finalTrees.GetTree(i);
                    writer.WriteLine("| {0,-5}| {1,-15}| {2,-10}| {3,-10}| {4,-10}", i + 1, tree.Type, tree.Age,
                                     tree.Height, tree.result);
                }
            }
        }
    }
}
