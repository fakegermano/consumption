#!/usr/bin/env python3
import csv
import math
from datetime import datetime

with open("meterusage.csv", "r", newline="") as meterusage_in:
    meter_reader = csv.reader(meterusage_in, delimiter=",", quotechar='"')
    with open("meterusage_rfc.csv", "w", newline="") as meterusage_out:
        meter_writer = csv.writer(
            meterusage_out, delimiter=",", quotechar='"', quoting=csv.QUOTE_NONE
        )
        for i, row in enumerate(meter_reader):
            if i == 0:
                meter_writer.writerow(row)
            else:
                date, measure = row

                date = datetime.fromisoformat(date)
                measure = float(measure)
                if math.isnan(measure):
                    measure = -1.0
                meter_writer.writerow([date.isoformat("T") + "Z", measure])
