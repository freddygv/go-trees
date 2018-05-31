library(dplyr)
library(ggplot2)
library(ggthemes)

df <- read.table(file = 'output.log', sep = '\t', header = FALSE)
colnames(df) <- c("operation", "input", "n", "data_structure", "rounds", "time")

agg <- group_by(df, operation, input, data_structure, n) %>%
    summarise(time = mean(time))

png("bench-time.png", width=1200, height=600, pointsize=18)

mutate(agg, time=time*1e-6) %>%

ggplot(aes(x=n, y=time, color=data_structure)) + 
    geom_line(size=1) + 
    facet_grid(operation ~ input) +
    xlab("Sample Size") + 
    ylab("Duration (ms)") +
    labs(color="Data Structure") +
    ggtitle("Time by Operation, Input Type, and Input Size") +
    theme_hc() +
    theme(title = element_text(size = 20)) +
    theme(legend.text = element_text(size = 16), legend.title=element_blank(), legend.position="right") +
    theme(axis.title.x = element_text(size = 18), axis.title.y = element_text(size = 18)) +
    theme(axis.text.x = element_text(size = 14), axis.text.y = element_text(size = 14)) +
    theme(strip.text.x = element_text(size = 14), strip.text.y = element_text(size = 14))

dev.off()